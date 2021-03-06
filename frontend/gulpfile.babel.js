const gulp = require('gulp'),
    sourcemaps = require('gulp-sourcemaps'),
    concat = require('gulp-concat'),
    sass = require('gulp-sass'),
    fs = require("fs"),
    colors = require('colors'),
    browserify = require('browserify'),
    source = require('vinyl-source-stream'),
    buffer = require('vinyl-buffer'),
    express = require('express'),
    bs = require('browser-sync').create(),
    reload = bs.reload,
    multiDest = require("gulp-multi-dest"),
    cleanCSS = require('gulp-clean-css'),
    shell = require('gulp-shell'),
    path = require('path'),
    postcss = require('gulp-postcss'),
    autoprefixer = require('autoprefixer');

const appRoot = path.resolve(__dirname);

const log = (o, level = 0) => {
    if (level > 2)
        return;
    for (var p in o) {
        console.log(`${colors.red('prop:')}${p}: ${o[p]}`);
        if (o[p] != null && typeof o[p] == 'object') {
            try {
                console.log("DETAILS");
                log(o[p], level + 1);
            } catch (err) {
                console.log('CANT GET INFO');
            }
        }
    }
};

let router = express.Router();
let jsonServer = require('json-server');
let server = null;

const config = {
    distribution: {
        js: ["../www/js","./dist/js"],
        html: ["../www","./dist"],
        css: ["../www/css","./dist/css"],
        images: ["../www/img","./dist/img"],
    }
};

var jsonData = require('./src/data/generate.js');

var packageJSON = require('./package.json');
var dependencies = Object.keys(packageJSON && packageJSON.dependencies || {});

const json = (callback) => {
    console.log(colors.cyan('[JSON] Generating a new DB'));

    // Must delete cached object and re-require or else
    delete require.cache[require.resolve('./src/data/generate.js')];
    jsonData = require('./src/data/generate.js');

    try {
        fs.writeFile("./src/data/db.json", JSON.stringify(jsonData()), 'utf8', (err) => {
            if (err) {
                console.log('[JSON] ' + colors.red(err));
                if (callback)
                    callback();
            } else {
                console.log(colors.green('[JSON] DB.json Saved'.bold));
                if (callback)
                    callback();
            }
        });
    } catch (err) {
        console.log('[JSON] ' + colors.red(err.toString()));
        if (callback)
            callback();
    }
};

const html = (callback) => {
    console.log('[HTML] ' + colors.cyan('Copying HTML'));
    return gulp.src('./src/markup/**/*.*')
        .pipe(multiDest(config.distribution.html))
        .on('error', (err) => {
            console.log('[HTML] ' + colors.bgWhite.red(err.toString()));
            callback();
        })
        .on('end', () => {
            callback();
        })
};

const img = (callback) => {
    console.log(colors.cyan('[IMAGE] Copying Images'));
    return gulp.src('./src/img/**/*.*')
        .pipe(multiDest(config.distribution.images))
        .on('error', function (err) {
            console.log('[IMAGE] ' + colors.red(err.toString()));
            callback();
        }).on('end', function () {
            callback();
        });
};

const jsbundle = (input, output, destinations, callback) => {
    var b = browserify({
        entries: input,
        debug: true
    })
        .transform('babelify', {
            presets: ["@babel/preset-env", "@babel/preset-react"]
        });

    return bundleJS(b, output, destinations, callback);
};

const bundleJS = (browserify, output, destinations, callback) => {
    return browserify
        .bundle((err) => {
            if (err)
                console.log('[JS] ' + colors.red(err.toString()));

            if (callback)
                callback();
        })
        .on('error', function (err) {
            console.log('[JS] ' + colors.red(err.toString()));
            callback();
        })
        .on('end', function () {
            callback();
        })
        .pipe(source(output))
        .pipe(buffer())
        .pipe(sourcemaps.init({
            loadMaps: true
        }))
        .pipe(sourcemaps.write('./'))
        .pipe(multiDest(destinations));
};

const js = (callback) => {
    console.log(colors.cyan('[JS] Bundling and Babeling JS'));
    jsbundle('./app.js', 'app.min.js', config.distribution.js, callback);
};

const jsv = (callback) => {
    console.log(colors.cyan('[JS V] Bundling and Babeling Vendor JS'));
    var b = browserify({
        debug: true
    }).transform('babelify', {
        presets: ['@babel/preset-env']
    });
    dependencies.forEach(lib => {
        b.require(lib);
    });

    return bundleJS(b, 'vendor.min.js', config.distribution.js, callback);
};

const scss = (callback) => {
    console.log(colors.cyan('[SCSS] Transpiling Global Sass to Css'));
    return bundleCSS([
        './src/styles/global.scss'
    ], 'bundle.min.css', callback);

};

function bundleCSS(source, dest, callback) {
    return gulp.src(source)
        .pipe(sourcemaps.init())
        .pipe(sass().on('error', sass.logError))
        .pipe(concat(dest))
        .pipe(postcss([autoprefixer()]))
        .pipe(cleanCSS({
            compatibility: 'ie11'
        }))
        .pipe(sourcemaps.write('.'))
        .on('end', callback)
        .on('error', function (err) {
            console.log(colors.red('[SCSS] ' + err.toString()));
            callback();
        })
        .pipe(multiDest(config.distribution.css))
        .pipe(bs.stream());

}

const serve = (callback) => {
    console.log(colors.cyan('[SERVE] Says: standing up your server'));
    inject_middleware();
    var app = express();
    app.use(['/discover*'], function (req, res) {
        res.sendFile(appRoot + './dist/index.html');
    });
    bs.init({
        open: false,
        https: true,
        notify: true,
        logPrefix: 'Server Says:',
        server: {
            baseDir: "./dist/",
            index: "index.html"
        },
        middleware: [
            function (req, res, next) {
                router(req, res, next);
            },
            app
        ]
    }, function () {
        console.log(colors.cyan('[SERVE] Says: hello'));
        callback();
    });
};

const inject_middleware = (cb) => {
    console.log(colors.cyan('[ROUTE] Rebuilding routes'));
    router = express.Router();
    server = jsonServer.create({
        verbosity: {
            level: "info",
            urlTracing: false
        }
    });
    server.use(jsonServer.defaults());
    server.use(jsonServer.router(jsonData()));
    router.use('/api/v1', server);
    if (cb) cb();
};

let WatchQueue = require('./lib/watch-queue');
let queue = new WatchQueue();
const watch = (done) => {

    console.log(colors.cyan('[WATCH] Watching...'));

    gulp.watch(['./src/markup/**/*.html'])
        .on('all', function () {
            queue.queue({
                name: 'HTML'
            }, (task, cb) => {
                bs.notify("Transpiling" + task.name, 1000);
                html(() => {
                    bs.notify("Done Transpiling" + task.name, 1000);
                    cb();
                });
            });
        });

    gulp.watch(['./src/**/*.scss'])
        .on('all', function () {
            queue.queue({
                name: 'Scss'
            }, (task) => {
                bs.notify("Transpiling" + task.name, 1000);
                scss(() => {
                    bs.notify("Done Transpiling Global " + task.name, 1000);
                });
            })
        });

    gulp.watch(['*.js', './src/*.js', './src/js/**/*.js', './src/components/**/*.js', './src/pages/**/*.js'])
        .on('all', function () {
            queue.queue({
                name: 'Js'
            }, (task) => {
                bs.notify("Transpiling" + task.name, 1000);
                js(() => {
                    bs.notify("Done Transpiling" + task.name, 1000);
                    reload();
                });
            });
        });

    gulp.watch(['./src/data/generate.js'])
        .on('all', function () {
            queue.queue({
                name: 'Generate'
            }, () => {
                bs.notify("Regenerating Data", 1000);
                json(() => {
                    inject_middleware(() => {
                        reload();
                        done();
                    });
                });
            });
        });

    gulp.watch(['./src/img/**/*'])
        .on('all', function () {
            queue.queue({
                name: 'Generate'
            }, () => {
                bs.notify("Transferring Images", 1000);
                img(() => {
                    reload();
                    done();
                });
            });
        });

    gulp.watch('./src/**/*')
        .on('all', function (event, path) {
            console.log(colors.yellow('File ' + path + ' ' + event));
        });

    done();
};

gulp.task('watch', watch);
gulp.task('build', gulp.series(gulp.parallel(html, scss, js, jsv, img)));
gulp.task('default', gulp.series(json, gulp.parallel(html, scss, js, jsv, img), gulp.parallel(serve, watch)));
gulp.task('serve', serve);
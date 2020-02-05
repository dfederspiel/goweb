/** Data generator script for mock service.
 *
 * Referenced during PUG transpilation to generate random content for static files
 * Consumed by json-server middleware (Browsersync and Node) to inject a mock API
 *
 * */

// https://lodash.com/
// https://github.com/Marak/faker.js

module.exports = function () {
    const faker = require('faker');
    const _ = require('lodash');
    return {
        pets: _.times(20, function (n) {
            const pet = {
                id: n,
                name: faker.name.firstName(),
                age: faker.random.number({min: 2, max: 18})
            };
            return pet;
        })
    };
};
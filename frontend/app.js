import React from 'react';
import ReactDOM from 'react-dom';
import {BrowserRouter as Router, Route} from "react-router-dom";

import Home from './src/components/home';

function App() {
    return (
        <Router>
            <div>
                <Route path="/">
                    <Home/>
                </Route>
            </div>
        </Router>
    )
}

ReactDOM.render(<App/>, document.getElementById('app'))
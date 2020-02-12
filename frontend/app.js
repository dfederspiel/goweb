import React, {useState, useEffect} from 'react';
import ReactDOM from 'react-dom';
import {BrowserRouter as Router, Route} from "react-router-dom";
import oauthSignIn, {getActiveUser, logout} from './src/js/auth'

function Home() {

    const [user, setUser] = useState("");

    useEffect(() => {
        getActiveUser(setUser)
    }, []);

    if(user === "")
        return <h1>Welcome! Please&nbsp;<a id="user-login" onClick={oauthSignIn}>log in</a>&nbsp;to continue </h1>
    else 
        return (
            <div>
                <h1>Hello, {user}&nbsp;<form method={"POST"} action="/logout"><button type="submit">Logout</button></form></h1>
            </div>
        )
}

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
import React, {useState} from 'react';
import ReactDOM from 'react-dom';
import {BrowserRouter as Router, Route} from "react-router-dom";
import oauthSignIn, {getActiveUser, logout} from './src/js/auth'

function Home() {

    const [user, setUser] = useState("")

    getActiveUser(setUser)

    if(user == "")
        return <div>Welcome! Please&nbsp;<a id="user-login" onClick={oauthSignIn}>log in</a>&nbsp;to continue </div>
    else 
        return <div>Hello, {user}<br /><form method="post" action="/logout"><button type="submit">Logout</button></form></div>
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
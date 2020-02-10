import React, {useState} from 'react';
import ReactDOM from 'react-dom';
import {BrowserRouter as Router, Route} from "react-router-dom";
import oauthSignIn, {getActiveUser} from './src/js/auth'

function Home() {

    const [user, setUser] = useState("")

    getActiveUser(setUser)

    if(user == "")
        return <div>Welcome! Please <a id="user-login" onClick={oauthSignIn}>log in</a> to continue </div>
    else 
        return <div>Hello, {user}</div>
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
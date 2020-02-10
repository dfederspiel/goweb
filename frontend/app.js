import React, {useState} from 'react';
import ReactDOM from 'react-dom';
import {BrowserRouter as Router, Route} from "react-router-dom";


/*
 * Create form to request access token from Google's OAuth 2.0 server.
 */
function oauthSignIn() {
    // Google's OAuth 2.0 endpoint for requesting an access token
    var oauth2Endpoint = 'https://accounts.google.com/o/oauth2/v2/auth';

    // Create <form> element to submit parameters to OAuth 2.0 endpoint.
    var form = document.createElement('form');
    form.setAttribute('method', 'GET'); // Send as a GET request.
    form.setAttribute('action', oauth2Endpoint);

    // Parameters to pass to OAuth 2.0 endpoint.
    var params = {
        'client_id': '90445840135-99mhv65o8m5kt3n6v46h6k1c2ie0eum1.apps.googleusercontent.com',
        'redirect_uri': 'http://localhost:8080/callback',
        'response_type': 'code',
        'scope': 'openid profile email',
        'include_granted_scopes': 'true',
        'state': ''
    };

    // Add form parameters as hidden input values.
    for (var p in params) {
        var input = document.createElement('input');
        input.setAttribute('type', 'hidden');
        input.setAttribute('name', p);
        input.setAttribute('value', params[p]);
        form.appendChild(input);
    }

    // Add form to page and submit it to open the OAuth 2.0 endpoint.
    document.body.appendChild(form);
    form.submit();
}

const getActiveUser = async (setUser) => {
    try {
        var response = await fetch("/api/v2/user")
        if(response.status != 401){
            var json = await response.json()
            console.log(json)
            setUser(json.email)
        }
    } catch {
        console.log("uh-oh")
    }
}

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
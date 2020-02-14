import React, {useState, useEffect} from 'react';
import oauthSignIn, {getActiveUser, logout} from '../js/auth'

export default function Home() {

    const [user, setUser] = useState("");

    useEffect(() => {
        getActiveUser(setUser)
    }, []);

    if(user === "")
        return <h1>Welcome! Please&nbsp;<a id="user-login" onClick={oauthSignIn}>log in</a>&nbsp;to continue </h1>
    else
        return (
            <div>
                <h1>Hello, {user}&nbsp;<form method={"POST"} action={"/logout"}><button type="submit">Logout</button></form></h1>
            </div>
        )
}
export default function oauthSignIn() {
    // Google's OAuth 2.0 endpoint for requesting an access token
    let oauth2Endpoint = 'https://accounts.google.com/o/oauth2/v2/auth';

    // Create <form> element to submit parameters to OAuth 2.0 endpoint.
    let form = document.createElement('form');
    form.setAttribute('method', 'GET'); // Send as a GET request.
    form.setAttribute('action', oauth2Endpoint);

    // Parameters to pass to OAuth 2.0 endpoint.
    let params = {
        'client_id': '90445840135-99mhv65o8m5kt3n6v46h6k1c2ie0eum1.apps.googleusercontent.com',
        'redirect_uri': 'http://localhost:8080/callback',
        'response_type': 'code',
        'scope': 'openid profile email',
        'include_granted_scopes': 'true',
        'state': ''
    };

    // Add form parameters as hidden input values.
    for (let p in params) {
        let input = document.createElement('input');
        input.setAttribute('type', 'hidden');
        input.setAttribute('name', p);
        input.setAttribute('value', params[p]);
        form.appendChild(input);
    }

    // Add form to page and submit it to open the OAuth 2.0 endpoint.
    document.body.appendChild(form);
    form.submit();
}

export const getActiveUser = (setUser) => {
    fetch("/api/v2/user").then(async r => {
        if (r.status !== 401) {
            let json = await r.json()
            setUser(json.name)
        }
    })

}

export const logout = async () => {
    //var response = await fetch("/logout")
    window.location.href = "/logout"
}
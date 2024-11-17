import { useState } from "react";

export function useAuth() {
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    return { isAuthenticated, setIsAuthenticated };
}


// TODO: Add authentication with the team server 
// / This auth should be similar to Starkiller's auth with the Empire server.
// / - User, Password, URL of team server
// / - Save the user's JWT in LocalStorage?



export function clientLogin() {
    // TODO: Implement login
    // / - should provide the user with a login prompt
    // / - this form should send an HTTP POST request to the team server's `/login` endpoint
    // / - if the login is successful, the user's JWT should be saved in LocalStorage
    // / - if the login is successful, the user should be redirected to the dashboard
    // / - if the login fails, the user should be notified with an error message
    console.log("Login");
}

export function clientLogout() {
    // TODO: Implement logout
    // / - should provide the user with a logout prompt
    // / - this form should send an HTTP POST request to the team server's `/logout` endpoint
    // / - if the logout is successful, the user's JWT should be removed from LocalStorage
    // / - if the logout fails, the user should be notified with an error message
}


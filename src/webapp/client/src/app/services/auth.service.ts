import { Injectable, Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Subject } from 'rxjs';
import { Router } from '@angular/router';
import { AppConfig } from '../config/api-config';

@Injectable({ providedIn: 'root' })
export class AuthenticationService {
    private _apiUrl = AppConfig.apiUrl;
    private _apiAuth = AppConfig.apiAuth;
    private _apiLogin = AppConfig.apiLogin;

    private _isAuthenticated = false;
    private _token: string;
    private _tokenTimer: any;
    private _authenticationStatusListener = new Subject<boolean>();

    constructor(
        private _http: HttpClient,
        private _router: Router,
    ) {}


    // Get the users authentication token
    getToken() {
        return this._token;
    }

    // Return if user is authenticated
    getIsAuthenticated() {
        return this._isAuthenticated;
    }

    // Return the authentication status listener as observable object
    getAuthStatusListener() {
        return this._authenticationStatusListener.asObservable();
    }

    // Login user
    login(email: string, name: string, image: string, id_token: string, userid: string) {
        const loginCredentials = {email, name, image, id_token, userid};
        this._http.post<{token: string, expiresIn: number, userid: string}>(this._apiUrl + this._apiAuth + this._apiLogin, loginCredentials)
            .subscribe(response => {
                console.log(response);
                const token = response.token;
                this._token = token;
                if (token) {
                    this._isAuthenticated = true;
                    this._authenticationStatusListener.next(true);
                    const expiresInDuration = response.expiresIn;
                    this.setAuthTimer(expiresInDuration);
                    const now = new Date();
                    const expirationDate = new Date(now.getTime() + expiresInDuration * 1000);
                    this.saveAuthenticationData(token, expirationDate, response.userid);
                    window.location.reload();
                    this._router.navigate(['/']);
                }
            });
    }

    // Automatically authenticate the user if they return to the site but session not expired
    autoAuthUser() {
        const authInformation = this.getAuthenticationData();
        if (!authInformation) {
            return;
        }
        const now = new Date();
        const expiresIn = authInformation.expirationDate.getTime() - now.getTime();
        if (expiresIn > 0) {
            this._token = authInformation.token;
            this._isAuthenticated = true;
            this.setAuthTimer(expiresIn / 1000);
            this._authenticationStatusListener.next(true);
        }
    }

    // Log the user out and end the session
    logout() {
        this._token = null;
        this._isAuthenticated = false;
        this._authenticationStatusListener.next(false);
        this._router.navigate(['/']);
        this.clearAuthenticationData();
        clearTimeout(this._tokenTimer);
    }


    // Set the authentication timer
    private setAuthTimer(duration: number) {
        this._tokenTimer = setTimeout(() => {
            this.logout();
        }, duration * 1000);
    }

    // Save authentication data to browser local storage
    private saveAuthenticationData(token: string, expirationDate: Date, id: string) {
        localStorage.setItem('token', token);
        localStorage.setItem('expiration', expirationDate.toISOString());
        localStorage.setItem('uid', id);
    }

    // Clear the users authentication data from browser local storage
    private clearAuthenticationData() {
        localStorage.removeItem('token');
        localStorage.removeItem('expiration');
        localStorage.removeItem('uid');
    }

    // Get the users authentication data from browser local storage
    private getAuthenticationData() {
        const token = localStorage.getItem('token');
        const expirationDate = localStorage.getItem('expiration');
        if (!token || !expirationDate) {
            return;
        }
        return {
            token: token,
            expirationDate: new Date(expirationDate)
        };
    }
}

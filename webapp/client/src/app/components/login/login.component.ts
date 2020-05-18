import { Component, OnInit } from '@angular/core';
import { AuthService, GoogleLoginProvider } from 'angular-6-social-login-v2';
import { AuthenticationService } from '../../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  providers: [AuthenticationService],
})

export class LoginComponent implements OnInit {

  constructor(
    private socialAuthService: AuthService,
    private authService: AuthenticationService,
    private router: Router,
  ) { }

  ngOnInit() {
  }

  // Login/Signup the user via Google OAuth API
  public socialLogin(socialPlatform: string) {
    let socialPlatformProvider;
    if (socialPlatform === 'google') {
      socialPlatformProvider = GoogleLoginProvider.PROVIDER_ID;
    }

    this.socialAuthService.signIn(socialPlatformProvider).then(
      (userData) => {
        this.authService.login(userData.email, userData.name, userData.image, userData.idToken, userData.id);
      }
    );
  }

}

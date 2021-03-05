import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { AuthService, AuthServiceConfig } from 'angular-6-social-login-v2';
import { getAuthServiceConfigs } from '../../app.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { LoginComponent } from './login.component';
import { RouterTestingModule } from '@angular/router/testing';

describe('LoginComponent', () => {
  let component: LoginComponent;
  let fixture: ComponentFixture<LoginComponent>;
  const config = getAuthServiceConfigs();
  beforeEach(async(() => {
    TestBed.configureTestingModule({
      providers: [AuthService, AuthServiceConfig],
      declarations: [ LoginComponent ],
      imports: [RouterTestingModule, HttpClientTestingModule]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  xit('should create', () => {
    expect(component).toBeTruthy();
  });
});

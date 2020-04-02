import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HttpModule } from '@angular/http';
import { HttpClientModule } from '@angular/common/http';
import cronstrue from 'cronstrue';
import { ChartsModule } from 'ng2-charts';



import {
  SocialLoginModule,
  AuthServiceConfig,
  GoogleLoginProvider
} from 'angular-6-social-login-v2';

import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { LoginComponent } from './components/login/login.component';
import { MainComponent } from './components/main/main.component';
import { SidebarComponent } from './components/sidebar/sidebar.component';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { CardBlockComponent } from './components/card-block/card-block.component';
import { CardJobInfoComponent } from './components/card-job-info/card-job-info.component';
import { CardJobRunsComponent } from './components/card-job-runs/card-job-runs.component';
import { CardJobStatusComponent } from './components/card-job-status/card-job-status.component';
import { CardBlockTwoComponent } from './components/card-block-two/card-block-two.component';
import { CardJobListComponent } from './components/card-job-list/card-job-list.component';
import { CardJobMetricsComponent } from './components/card-job-metrics/card-job-metrics.component';
import { MetricsLineChartComponent } from './components/metrics-line-chart/metrics-line-chart.component';
import { CardJobValuesListComponent } from './components/card-job-values-list/card-job-values-list.component';
import { AppConfig } from './config/api-config';


export function getAuthServiceConfigs() {
  const config = new AuthServiceConfig(
    [
      {
        id: GoogleLoginProvider.PROVIDER_ID,
        provider: new GoogleLoginProvider('150239218178-72l3odnt9e0ldo2dkpmqcj00sg6j2aun.apps.googleusercontent.com')
      }
    ]
  );
  return config;
}

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    MainComponent,
    SidebarComponent,
    CardBlockComponent,
    CardJobInfoComponent,
    CardJobRunsComponent,
    CardJobStatusComponent,
    CardBlockTwoComponent,
    CardJobListComponent,
    CardJobMetricsComponent,
    MetricsLineChartComponent,
    CardJobValuesListComponent
  ],
  imports: [
    SocialLoginModule,
    BrowserModule,
    HttpModule,
    HttpClientModule,
    AppRoutingModule,
    FontAwesomeModule,
    ChartsModule
  ],
  providers: [{
    provide: AuthServiceConfig,
    useFactory: getAuthServiceConfigs
  }],
  bootstrap: [AppComponent]
})
export class AppModule { }

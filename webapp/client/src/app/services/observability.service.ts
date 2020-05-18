import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { AppConfig } from '../config/api-config';
import { AuthenticationService } from '../services/auth.service';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ObservabilityService {
  private _apiUrl = AppConfig.apiUrl;
  private _apiObservability = AppConfig.apiObservability;
  private _apiJobRunStatus = AppConfig.apiJobRunStatus;
  private _apiJobMetrics = AppConfig.apiJobMetrics;

  public _StringMetricsSubject = new BehaviorSubject<any>([]);
  private _job: any = {};

  constructor(    
    private _http: HttpClient,
    private _router: Router,
    private _authService: AuthenticationService) { }

    setStringMetrics(StringMetrics: any) {
      this._StringMetricsSubject.next(StringMetrics);
    }

    jobRunStatus(jobId: String) {
      const token = this._authService.getToken()
      return this._http.post<{success: boolean, message: any}>(this._apiUrl + this._apiObservability + this._apiJobRunStatus, { token, jobId })
    }

    jobMetrics(jobId: String) {
      const token = this._authService.getToken()
      return this._http.post<{success: boolean, message: any}>(this._apiUrl + this._apiObservability + this._apiJobMetrics, { token, jobId })
    }
}

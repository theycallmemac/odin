import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { BehaviorSubject } from 'rxjs';
import { AppConfig } from '../config/api-config';
import { AuthenticationService } from '../services/auth.service';

@Injectable({
  providedIn: 'root'
})
export class JobsService {
  private _apiUrl = AppConfig.apiUrl;
  private _apiJobs = AppConfig.apiJobs;
  private _apiJobsGet = AppConfig.apiJobsGet;
  private _apiJobGet = AppConfig.apiJobGet;

  private _joba = {
    description: "Description",
    file: "file",
    gid: "gid",
    id: "id",
    language: "python3",
    name: "Name",
    runs: 1,
    schedule: "15 * * * *",
    stats: "",
    uid: "uid",
    _id: "id"
  }

  public _selectedJobListener = new BehaviorSubject<any>(this._joba);
  private _job: any = {};
  // private _authtoken: string;

  constructor(
    private _http: HttpClient,
    private _router: Router,
    private _authService: AuthenticationService
  ) { }


  // Return the selected job
  getSelectedJob() {
    return this._job
  }

  // Return the selected job listener as observable object
  getJobListener() {
    return this._selectedJobListener.asObservable();
  }

  // broadcasts the passed job via observable to subscribed objects and keeps a local copy
  setJob(job: any) {
    this._selectedJobListener.next(job);
    this._job = job;
  }

  // returns all jobs from the database
  jobsGet() {
    const token = this._authService.getToken()
    return this._http.post<{success: boolean, jobs: any}>(this._apiUrl + this._apiJobs + this._apiJobsGet, {token})
  }

  // returns a job by id
  jobGet() {
    const token = this._authService.getToken()
    this._http.post<{success: boolean, job: any}>(this._apiUrl + this._apiJobs + this._apiJobGet, { token })
        .subscribe(response => {
            return response.job
        });
  }
}
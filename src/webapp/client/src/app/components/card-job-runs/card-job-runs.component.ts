import { Component, OnInit } from '@angular/core';
import { JobsService } from '../../services/jobs.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-card-job-runs',
  templateUrl: './card-job-runs.component.html',
  styleUrls: ['./card-job-runs.component.css']
})
export class CardJobRunsComponent implements OnInit {
  private selectedJob: any;
  private jobSchedule: String;
  private selectedJobObv : any;
  
  constructor(
    private jobsService: JobsService
  ) { }

  ngOnInit() {
    this.selectedJobObv = this.jobsService._selectedJobListener
    this.selectedJobObv.subscribe(job => {
      this.selectedJob = job;
    });
  }

}

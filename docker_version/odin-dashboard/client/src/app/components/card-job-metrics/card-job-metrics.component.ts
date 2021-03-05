import { Component, OnInit } from '@angular/core';
import { JobsService } from '../../services/jobs.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-card-job-metrics',
  templateUrl: './card-job-metrics.component.html',
  styleUrls: ['./card-job-metrics.component.css']
})
export class CardJobMetricsComponent implements OnInit {

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

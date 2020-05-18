import { Component, OnInit } from '@angular/core';
import { JobsService } from '../../services/jobs.service';
import { ObservabilityService } from '../../services/observability.service';
import { Subscription } from 'rxjs';


@Component({
  selector: 'app-card-job-status',
  templateUrl: './card-job-status.component.html',
  styleUrls: ['./card-job-status.component.css']
})
export class CardJobStatusComponent implements OnInit {

  private selectedJob: any;
  private jobSchedule: String;
  private selectedJobObv: any;
  public jobLastRunStatus: String;
  public jobLastRunTime: any;
  
  constructor(
    private jobsService: JobsService,
    private observabilityService: ObservabilityService
  ) { }

  ngOnInit() {
    this.selectedJobObv = this.jobsService._selectedJobListener
    this.selectedJobObv.subscribe(job => {
      this.selectedJob = job;
      this.observabilityService.jobRunStatus(this.selectedJob.id).subscribe((response) => {
        try {
          // translate status code to string 
          if (response.message.status[0] == "2") {
            this.jobLastRunStatus = "Success";
          } else if (response.message.status[0] == "4"){
            this.jobLastRunStatus = "Warning";
          } else if (response.message.status[0] == "5"){
            this.jobLastRunStatus = "Failure";
          } else {
            this.jobLastRunStatus = response.message;
          }
          // convert run timestamp to human friendly format
          this.jobLastRunTime = new Date((Number(response.message.time)*1000));
        } catch(err) {
          this.jobLastRunStatus = "No job results found"
          this.jobLastRunTime =  "No job results found"
        }
      })
    })
  }

}

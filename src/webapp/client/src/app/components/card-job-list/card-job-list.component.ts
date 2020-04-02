import { Component, OnInit } from '@angular/core';
import { JobsService } from '../../services/jobs.service';

@Component({
  selector: 'app-card-job-list',
  templateUrl: './card-job-list.component.html',
  styleUrls: ['./card-job-list.component.css'],
})
export class CardJobListComponent implements OnInit {

  
  message:string;

  public jobsFound: boolean = false;
  public jobList: any;
  public selectedJobObv;
  public selectedJob; 

  constructor(
    private jobsService: JobsService,
  ) { }

  ngOnInit() {
      this.jobsService.jobsGet().subscribe((response) => {
      this.jobList = response
      this.selectJob(response[3])
      this.selectedJob = response[3]
      this.jobsFound = true;
    });
      this.selectedJobObv = this.jobsService._selectedJobListener
      this.selectedJobObv.subscribe(job => {
      this.selectedJob = job; 
    })
  }

  // converts date within mongoDB _id string to a JS Date object
  // input : mongoDB _id (String)
  // output : timestamp (Date)
  dateFromObjectId(objectId: String) {
    return new Date(parseInt(objectId.substring(0, 8), 16) * 1000);
  };
  
  // get time in days since given Date
  // input : timestamp (Date)
  // output : Days since timestamp (num)
  dateDifference(first) {
    let date: any = new Date();
    return Math.round((date - first)/(1000*60*60*24));
  };

  selectJob(job: any) {
    // console.log(job)
    this.jobsService.setJob(job)
  }
}

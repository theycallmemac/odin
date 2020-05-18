import { Component, OnInit } from '@angular/core';
import { JobsService } from '../../services/jobs.service';
import { Subscription } from 'rxjs';
import cronstrue from 'cronstrue';

@Component({
  selector: 'app-card-job-info',
  templateUrl: './card-job-info.component.html',
  styleUrls: ['./card-job-info.component.css'],
})
export class CardJobInfoComponent implements OnInit {

  public selectedJob: any;
  public jobSchedule: String;
  public selectedJobObv : any;
  
  constructor(
    private jobsService: JobsService
  ) { }

  ngOnInit() {
    this.selectedJobObv = this.jobsService._selectedJobListener
    this.selectedJobObv.subscribe(job => {
      this.selectedJob = job;
      try {
        this.jobSchedule = this.convertToReadaleSchedules(job.schedule);
      }
      catch(err) {
        console.log(err)
      }
    });
  }

  // converts job schedule(s) to human readable form
  // Input : schedule(s) String
  // Output : human readable schedule(s) String
  convertToReadaleSchedules(schedule: String) {
    let humanSchedules : Array<String> = [];
    schedule = this.strip(",", schedule);

    // check if multiple schedules are assigned to job
    if (schedule.includes(",")) {
      let schedules = schedule.split(",");
      // convert each schedule to human readable schedule
      for (let schedule of schedules) {
        humanSchedules.push(String(cronstrue.toString(String(schedule))));
      }

    } else {
      humanSchedules.push(String(cronstrue.toString(String(schedule))));
    }

    return humanSchedules.join(" & ")
  }

  // strips given characters from the end of the string
  // Input : s characters to strip (String), string to strip from (String)
  // Return : if succesful stripped string (String) otherwise unstripped string
  strip(s: String, string: String) {
    try {
      return string.replace(new RegExp("[" + s + "]+$"), "");
    } catch(err) {
      console.log(err)
      return string
    }
  }

}

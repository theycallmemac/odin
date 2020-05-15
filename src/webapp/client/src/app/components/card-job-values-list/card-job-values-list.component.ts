import { Component, OnInit } from '@angular/core';
import { ObservabilityService } from '../../services/observability.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-card-job-values-list',
  templateUrl: './card-job-values-list.component.html',
  styleUrls: ['./card-job-values-list.component.css']
})
export class CardJobValuesListComponent implements OnInit {

  private ValuesListObv : any;
  private ValuesList : any;

  public tableHeaders = ["timestamp", "type", "desc", "value"];
  public mappings = {
    'timestamp' : 'Run Time',
    'type' : 'type',
    'desc' : 'description',
    'value' : 'value'
    }


  constructor(
    private observablityService: ObservabilityService,
  ) { }

  ngOnInit() {
    this.ValuesListObv = this.observablityService._StringMetricsSubject
    this.ValuesListObv.subscribe(valuesList => {
      // convert unix dates to human friendly format
      for (let i = 0; i < valuesList.length; i++) {
        let date = new Date((Number(valuesList[i].timestamp)*1000))
        valuesList[i].timestamp = date.toLocaleString()
      }
      this.ValuesList = valuesList
    });
  }

}

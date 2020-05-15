import { Component, OnInit } from '@angular/core';
import { JobsService } from '../../services/jobs.service';
import { ObservabilityService } from '../../services/observability.service';

@Component({
  selector: 'app-metrics-line-chart',
  templateUrl: './metrics-line-chart.component.html',
  styleUrls: ['./metrics-line-chart.component.css']
})
export class MetricsLineChartComponent implements OnInit {

  constructor(
    private jobsService: JobsService,
    private observabilityService: ObservabilityService
    ) { }
  
  private selectedJob: any;
  private selectedJobObv: any;

  // Chart
  public chartloaded = false;
  public chartTitle;
  public lineChartType = 'line';
  public lineChartData: Array<any> = [];
  public lineChartLabels: Array<any> = [];
  public lineChartOptions: any = { responsive: true, line : {lineTension: 0 }};
  public lineChartLegend = true;
  public lineChartPlugins = [];


  ngOnInit() {
    try {
      this.selectedJobObv = this.jobsService._selectedJobListener
      this.selectedJobObv.subscribe(job => {
        this.chartloaded = false;
        this.selectedJob = job;
        this.observabilityService.jobMetrics(job.id).subscribe((response) => {
          if (response.success == true) {
            this.chartloaded = false;
            try {
            this.lineChartLabels = this.setLabels(response.message[0])
            let data = this.getData(response.message)
            this.lineChartData = data[0]
            this.observabilityService.setStringMetrics(data[1])
            this.chartloaded = true;
            } catch(err) {
              this.chartloaded = false;
            }
          }
        })
      })
    this.updater()
    } catch(err) {}
  }

  updater() {
    let _this = this;
    
    setInterval(function() {
      _this.observabilityService.jobMetrics(_this.selectedJob.id).subscribe((response) => {
        try {
          _this.lineChartLabels = _this.setLabels(response.message[0])
          let data = _this.getData(response.message)
          _this.lineChartData = data[0]
          _this.observabilityService.setStringMetrics(data[1])
          _this.chartloaded = true;
        } catch(err) {
          _this.chartloaded = false;
        }
      }
    )}, 3000)
  }

  getData(metrics: any) {
    let stringDataArray = [];
    let dataArray = [];

    for (let i = 0; i < metrics.length; i++) {
      metrics[i].observability = metrics[i].observability.slice(metrics[i].observability.length - 10, metrics[i].observability.length);
      if (metrics[i].observability[0].type !== "result" && metrics[i].observability[0].type !== "condition"  && !isNaN(Number(metrics[i].observability[0].value))) {
        let data =  { data: [], label: metrics[i]["_id"]}
        for (let j = 0; j < metrics[i].observability.length; j++) {
          data.data.push(Number(metrics[i].observability[j].value));
        }
        dataArray.push(data);
      } else if (metrics[i].observability[0].type == "condition") {
        let data =  { data: [], label: metrics[i]["_id"], lineTension: 0, fill: false }
        for (let j = 0; j < metrics[i].observability.length; j++) {
          if (metrics[i].observability[j].value.toLowerCase() == "true") {
            data.data.push(100);
          } else {
            data.data.push(0);
          }
        }
        dataArray.push(data);
      } else {
        for (let j = 0; j < metrics[i].observability.length; j++) {
          stringDataArray.push(metrics[i].observability[j])
        }
      }

    }
    if (dataArray !== []) {return [dataArray, stringDataArray];}
  }

  setLabels(metrics: any) {
    let labels = []
    const monthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun",
      "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];

    for (let i = 0; i < metrics.observability.length && i < 10 ; i++) {
      let date = new Date((Number(metrics.observability[i].timestamp)*1000))
      let day = date.getDate();
      let month = monthNames[date.getMonth()]
      let hours;
      if ( Number(date.getHours()) < 10) {
        hours = "0"+ String(date.getHours())
      } else { 
        hours = date.getHours() 
      }
      let minutes; 
      if ( Number(date.getMinutes()) < 10) {
        minutes = "0"+ String(date.getMinutes() )
      } else { 
        minutes = date.getMinutes() 
      }
      labels.push(day + " " + month + " " + hours + ":" + minutes);
    }
    return labels
  }
}

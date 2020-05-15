import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MetricsLineChartComponent } from './metrics-line-chart.component';

describe('MetricsLineChartComponent', () => {
  let component: MetricsLineChartComponent;
  let fixture: ComponentFixture<MetricsLineChartComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MetricsLineChartComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MetricsLineChartComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

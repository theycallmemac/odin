import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CardJobMetricsComponent } from './card-job-metrics.component';

describe('CardJobMetricsComponent', () => {
  let component: CardJobMetricsComponent;
  let fixture: ComponentFixture<CardJobMetricsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CardJobMetricsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CardJobMetricsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

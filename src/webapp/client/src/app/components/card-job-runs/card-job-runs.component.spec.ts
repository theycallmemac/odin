import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CardJobRunsComponent } from './card-job-runs.component';

describe('CardJobRunsComponent', () => {
  let component: CardJobRunsComponent;
  let fixture: ComponentFixture<CardJobRunsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CardJobRunsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CardJobRunsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

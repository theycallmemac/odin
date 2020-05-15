import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CardJobStatusComponent } from './card-job-status.component';

describe('CardJobStatusComponent', () => {
  let component: CardJobStatusComponent;
  let fixture: ComponentFixture<CardJobStatusComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CardJobStatusComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CardJobStatusComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

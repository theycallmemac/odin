import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CardJobInfoComponent } from './card-job-info.component';

describe('CardJobInfoComponent', () => {
  let component: CardJobInfoComponent;
  let fixture: ComponentFixture<CardJobInfoComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CardJobInfoComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CardJobInfoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

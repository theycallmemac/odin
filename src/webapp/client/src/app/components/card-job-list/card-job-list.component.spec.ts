import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CardJobListComponent } from './card-job-list.component';

describe('CardJobListComponent', () => {
  let component: CardJobListComponent;
  let fixture: ComponentFixture<CardJobListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CardJobListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CardJobListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

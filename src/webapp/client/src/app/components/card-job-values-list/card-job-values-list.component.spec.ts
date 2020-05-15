import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CardJobValuesListComponent } from './card-job-values-list.component';

describe('CardJobValuesListComponent', () => {
  let component: CardJobValuesListComponent;
  let fixture: ComponentFixture<CardJobValuesListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CardJobValuesListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CardJobValuesListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

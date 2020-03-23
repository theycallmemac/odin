import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CardBlockTwoComponent } from './card-block-two.component';

describe('CardBlockTwoComponent', () => {
  let component: CardBlockTwoComponent;
  let fixture: ComponentFixture<CardBlockTwoComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CardBlockTwoComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CardBlockTwoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

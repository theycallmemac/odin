import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { CardJobListComponent } from './card-job-list.component';
import { HttpClientTestingModule} from '@angular/common/http/testing';
import { RouterTestingModule } from '@angular/router/testing';

describe('CardJobListComponent', () => {
  let component: CardJobListComponent;
  let fixture: ComponentFixture<CardJobListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CardJobListComponent ],
      imports: [RouterTestingModule, HttpClientTestingModule]
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

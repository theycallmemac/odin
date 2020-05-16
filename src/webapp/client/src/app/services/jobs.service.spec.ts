import { TestBed, inject } from '@angular/core/testing';
import { HttpClientTestingModule} from '@angular/common/http/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { JobsService } from './jobs.service';

describe('JobsService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [JobsService],
      imports: [RouterTestingModule, HttpClientTestingModule]
    });
  });

  it('should be created', inject([JobsService], (service: JobsService) => {
    expect(service).toBeTruthy();
  }));
});

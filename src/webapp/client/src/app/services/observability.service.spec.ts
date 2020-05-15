import { TestBed, inject } from '@angular/core/testing';
import { HttpClientTestingModule} from '@angular/common/http/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { ObservabilityService } from './observability.service';

describe('ObservabilityService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [ObservabilityService],
      imports: [RouterTestingModule, HttpClientTestingModule]
    });
  });

  it('should be created', inject([ObservabilityService], (service: ObservabilityService) => {
    expect(service).toBeTruthy();
  }));
});

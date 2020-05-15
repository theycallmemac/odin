import { TestBed, inject } from '@angular/core/testing';

import { ObservabilityService } from './observability.service';

describe('ObservabilityService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [ObservabilityService]
    });
  });

  it('should be created', inject([ObservabilityService], (service: ObservabilityService) => {
    expect(service).toBeTruthy();
  }));
});

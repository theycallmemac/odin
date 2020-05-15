import { TestBed, inject } from '@angular/core/testing';

import { JobsService } from './jobs.service';

describe('JobsService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [JobsService]
    });
  });

  it('should be created', inject([JobsService], (service: JobsService) => {
    expect(service).toBeTruthy();
  }));
});

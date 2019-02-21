import { Component, OnInit } from '@angular/core';
import { ContextService } from '../context.service';

@Component({
  selector: 'app-listing',
  templateUrl: './listing.component.html',
  styleUrls: ['./listing.component.scss']
})
export class ListingComponent implements OnInit {

  constructor(private contextSvc: ContextService) { }

  ngOnInit() {
    this.getListing();
  }

  getListing(): void {
    this.listing = this.contextSvc.getListing();
  }
}

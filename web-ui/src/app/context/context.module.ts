import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule }    from '@angular/common/http';

import { ContextRoutingModule } from './context-routing.module';
import { ListingComponent } from './listing/listing.component';

@NgModule({
  declarations: [ListingComponent],
  imports: [
    CommonModule,
    HttpClientModule,
    ContextRoutingModule
  ]
})
export class ContextModule { }

import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { TableComponent } from "./table/table.component"

const routes: Routes = [
  {
    path:"table",
    component: TableComponent},
  {
    path: "context",
    loadChildren: "./context/context.module#ContextModule"
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

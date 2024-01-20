import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PublicComponent } from './public.component';
import { RegisterComponent } from './register/register.component';
import { LoginComponent } from './login/login.component';
import { FormsModule,ReactiveFormsModule } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { RouterOutlet } from '@angular/router';



@NgModule({
  declarations: [
    PublicComponent,
    RegisterComponent,
    LoginComponent
  ],
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClient,
    RouterOutlet
  ]
})
export class PublicModule { }

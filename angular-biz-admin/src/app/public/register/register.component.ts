import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
@Component({
  selector: 'app-register',
  standalone: true,
  imports: [],
  templateUrl: './register.component.html',
  styleUrl: './register.component.css'
})
export class RegisterComponent {
  firstName = '';
  lastName = '';
  email = '';
  password = '';
  passwordConfirm = '';
  //role = '';

  constructor(
    private http: HttpClient,
    private router: Router
  ){}

  submit(): void{
    console.log({
      first_name: this.firstName,
      last_name: this.lastName,
      email: this.email,
      password: this.password,
      password_confirm: this.passwordConfirm,
    });
    this.http.post('http://localhost:8080/api/register',{
      first_name: this.firstName,
      last_name: this.lastName,
      email: this.email,
      password: this.password,
      password_confirm: this.passwordConfirm,
    }).subscribe(() => {
      //login
      this.router.navigate(['/login'])
    })
    //处理异步
  }

}

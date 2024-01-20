import { Component } from '@angular/core';
import { FormGroup, FormBuilder } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {
  form!: FormGroup; // !前不可能为空，不需要非空检查
  constructor(private formBuilder: FormBuilder, private http: HttpClient, private router: Router){};
  // init __init__
  ngOnInit(): void{
    this.form = this.formBuilder.group({
      email: '',
      password:''
    });
  }

  submit():void {
    
    console.log(this.form.getRawValue())
    this.http.post('http://localhost:8080/api/login',this.form.getRawValue()).subscribe(
      () => {
        //首页
        this.router.navigate(['/'])
      }
    )
  }
}

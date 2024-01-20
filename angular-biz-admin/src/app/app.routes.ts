import { RouterModule, Routes } from '@angular/router';
import { PublicComponent } from './public/public.component';
import { LoginComponent } from './public/login/login.component';
import { RegisterComponent } from './public/register/register.component';
import { SecureComponent } from './secure/secure.component';
import { NgModel } from '@angular/forms';

export const routes: Routes = [
    {
        path: '',
        component: PublicComponent,
        children:[
            {path: 'login',component:LoginComponent},
            {path: 'register', component:RegisterComponent},
        ]
    },
    {
        path:'',
        component:SecureComponent,
        children:[
        ]
    }
];


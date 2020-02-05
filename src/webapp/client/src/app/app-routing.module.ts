import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './guards/auth.guard';
import { NoAuthGuard } from './guards/no-auth.guard';
import { MainComponent} from './components/main/main.component';
import { LoginComponent } from './components/login/login.component';

const appRoutes: Routes = [
    { path: '', component: MainComponent, canActivate: [AuthGuard]},
    { path: 'login', component: LoginComponent, canActivate: [NoAuthGuard] },
];

@NgModule({
    imports: [
        RouterModule.forRoot(
            appRoutes,
        )
    ],
    exports: [
        RouterModule
    ],
    providers: [
        AuthGuard,
        NoAuthGuard
    ]
})

export class AppRoutingModule { }

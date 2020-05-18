import { Component, OnInit } from '@angular/core';
import { faHome, faCog, faPowerOff  } from '@fortawesome/free-solid-svg-icons';
import { AuthenticationService } from '../../services/auth.service';


@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css']
})
export class SidebarComponent implements OnInit {
  faHome;
  faCog;
  faPowerOff;

  constructor(
    private authService : AuthenticationService
  ) { }

  ngOnInit() {
    this.faHome = faHome;
    this.faCog = faCog;
    this.faPowerOff = faPowerOff;
  }

  logout() {
    this.authService.logout()
  }
}

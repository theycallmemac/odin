import { Component, OnInit } from '@angular/core';
import { faHome, faCog, faPowerOff  } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css']
})
export class SidebarComponent implements OnInit {
  faHome;
  faCog;
  faPowerOff;

  constructor() { }

  ngOnInit() {
    this.faHome = faHome;
    this.faCog = faCog;
    this.faPowerOff = faPowerOff;
  }

}

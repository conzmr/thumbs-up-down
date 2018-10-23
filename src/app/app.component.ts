import { Component, OnInit } from '@angular/core';
import { Http, Response } from '@angular/http'
import { map } from 'rxjs/operators';


@Component({
  selector: 'app-root',
  templateUrl: '/web/indexApp.html',
  styleUrls: ['/css/index.css']
})

export class AppComponent  {

  title = 'thumbs-up-down';
  numComments: any = 15;
  numLikes: any = 10;
  showComments = false;

  c = ['hola', 'hola2','hola3','hola4','hola5'];

  showCommentsClick(){
    this.showComments = !this.showComments;
    console.log(this.showComments)
  }

  addComment(value){
    console.log(value)
    if (value !== '') {
      this.c.push(value);
      this.numComments += 1;
    }
  }
  
  // constructor(public http: Http){
  //   this.numComments = this.c.length;
  //   this.numLikes = 15;
  // }

  // ngOnInit(){
  //   this.http.get('http://localhost:8080/posts')
  //     .pipe(map((res: Response) => res.json()))
  //     .subscribe(data => this.c = data);
  // }
}

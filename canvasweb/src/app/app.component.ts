import { Component } from '@angular/core';
var lineset: LineInter[] = [];

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  title = 'Ink Canvas';

  randomWholeNum() {
    var ret = Math.random()*1024;
    if(ret>0 && ret<1024) {
      return ret;
    }else{
      this.randomWholeNum();
    }
  }

  drawLine(p1:Point, p2:Point){
    var canvas :any = document.getElementById("canvas");
    let n1 = this.randomWholeNum();
    let n2 = this.randomWholeNum();
    p1 = {x: n1, y:n2}
    let n3 = this.randomWholeNum();
    let n4 = this.randomWholeNum();
    p2 = {x:n3, y:n4}
    var l : LineInter;
    l = {point1:p1,point2:p2};
    for(let i=0 ; i<= lineset.length-1 ;i++){
      let temp:LineInter;
      temp = lineset[i];
      if(this.overlap_detector(l,temp)){
        return false;
      }
    }
    if (canvas.getContext) {
      var ctx=canvas.getContext("2d");
      ctx.beginPath();
      ctx.moveTo(n1,n2);
      ctx.lineTo(n3,n4);
      ctx.strokeStyle="black";
      ctx.lineWidth=4;
      ctx.stroke();
      this.drawPoint(p1.x,p1.y);
      this.drawPoint(p2.x,p2.y);
    }
    lineset.push(l);
  }
  drawPoint(x,y){
    var pointSize = 6;
    var canvas :any = document.getElementById("canvas");
    var ctx = canvas.getContext("2d");
    ctx.fillStyle = "#ff2626";
    ctx.beginPath();
    ctx.arc(x, y, pointSize, 0, Math.PI * 2, true);
    ctx.fill();
}
  overlap_detector(l1:LineInter ,l2:LineInter){
    var point1 = l1.point1
    var point2 = l1.point2
    var linePoint1 = l2.point1
    var linePoint2 = l2.point2
    var denominator = (point1.y-point2.y)*(linePoint1.x-linePoint2.x)-(point2.x-point1.x)*(linePoint2.y-linePoint1.y);
    if(denominator==0) {
      return false;
    }
    var x = ((point1.x-point2.x)*(linePoint1.x-linePoint2.x)*(linePoint2.y-point2.y)+
      (point1.y-point2.y)*(linePoint1.x-linePoint2.x)*point2.x-(linePoint1.y-linePoint2.y)*
      (point1.x-point2.x)*linePoint2.x)/denominator;
    var y = -((point1.y-point2.y)*(linePoint1.y-linePoint2.y)*(linePoint2.x-point2.x)+
      (point1.x-point2.x)*(linePoint1.y-linePoint2.y)*point2.y-(linePoint1.x-linePoint2.x)*
      (point1.y-point2.y)*linePoint2.y)/denominator;
    if((x-point2.x)*(x-point1.x)<=0 && (y-point2.y)*(y-point1.y)<=0 && (x-linePoint2.x)*(x-linePoint1.x)<=0 &&
      (y-linePoint2.y)*(y-linePoint1.y)<=0) {
      return true;
    }
    return false;
  }

}
export interface LineInter {
  point1:     Point;
  point2:     Point;
}
export interface Point {
  x:     number;
  y:     number;
}

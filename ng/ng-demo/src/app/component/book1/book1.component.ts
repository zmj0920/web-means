import { addBook } from '@store/actions';
import { AppStoreModule } from '@store/store.module';
import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
@Component({
  selector: 'app-book1',
  templateUrl: './book1.component.html',
  styleUrls: ['./book1.component.css']
})
export class Book1Component implements OnInit {
  public bookName: string;
  public author: string;
  public price: number;

  constructor (private store: Store<AppStoreModule>) { }

  ngOnInit() {
   
  }

  public submit(): void {
    if (this.bookName && this.author && this.price) {
      this.store.dispatch(addBook({ book: { bookName: this.bookName, author: this.author, price: this.price, createAt: Date.now() } }));
      this.bookName = '';
      this.author = '';
      this.price = null;
    }
  }
}

import { addBook } from '@store/actions';
import { AppStoreModule } from '@store/store.module';
import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder } from '@angular/forms';
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
  profileForm = this.fb.group({
    bookName: [''],
    author: [''],
    price: [''],
    createAt: Date.now()
  });
  constructor(private fb: FormBuilder, private store: Store<AppStoreModule>) { }

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

  onSubmit() {
    // TODO: Use EventEmitter with form value
    console.warn(this.profileForm.value);
    this.store.dispatch(addBook({ book: { ...this.profileForm.value } }));
  }
}

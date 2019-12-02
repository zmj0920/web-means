import { addBook } from '@store/actions';
import { AppStoreModule } from '@store/store.module';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Store } from '@ngrx/store';
import { Subscription, of} from 'rxjs';
import { filter, map } from 'rxjs/operators';

interface Lengthwise {
  num: number;
}
@Component({
  selector: 'app-book1',
  templateUrl: './book1.component.html',
  styleUrls: ['./book1.component.css']
})
export class Book1Component implements OnInit {
  public bookName: string;
  public author: string;
  public price: number;
  public name: Lengthwise;
  profileForm = this.fb.group({
    bookName: ['', Validators.required],
    author: ['', Validators.required],
    price: ['', Validators.required],
    createAt: Date.now()
  });
  private subscription: Subscription; // 订阅
  constructor(private fb: FormBuilder, private store: Store<AppStoreModule>) { }

 someValue = 123;

 strLength: number = (this.someValue as unknown  as string).length;


  loggingIdentity<T extends Lengthwise>(arg: T[]): T[] {
    console.log(arg.length);
    return arg;
  }

  ngOnInit() {
    // this.loggingIdentity({ length: 10, value: 3 });
    this.subscription = of('桃子', '鲤鱼').pipe(
      filter(v => v === '桃子'),
      map(v => v + '罐头')
    ).subscribe(console.log);


    this.profileForm.valueChanges.pipe(
      filter(() => this.profileForm.valid) ,
      map(data => {
        data.lastTime = new Date();
        return data;
      })
    ).subscribe(res => console.log(res));
  }

  // tslint:disable-next-line: use-lifecycle-interface
  ngOnDestroy() {
    // 取消订阅
    this.subscription.unsubscribe();
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

  updateProfile() {
    this.profileForm.patchValue({
      bookName: 'Nancy'
    });
  }
}

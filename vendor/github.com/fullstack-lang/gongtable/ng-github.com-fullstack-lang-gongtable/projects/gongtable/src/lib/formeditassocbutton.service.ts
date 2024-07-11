// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { FormEditAssocButtonAPI } from './formeditassocbutton-api'
import { FormEditAssocButton, CopyFormEditAssocButtonToFormEditAssocButtonAPI } from './formeditassocbutton'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class FormEditAssocButtonService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  FormEditAssocButtonServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private formeditassocbuttonsUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.formeditassocbuttonsUrl = origin + '/api/github.com/fullstack-lang/gongtable/go/v1/formeditassocbuttons';
  }

  /** GET formeditassocbuttons from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormEditAssocButtonAPI[]> {
    return this.getFormEditAssocButtons(GONG__StackPath, frontRepo)
  }
  getFormEditAssocButtons(GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormEditAssocButtonAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<FormEditAssocButtonAPI[]>(this.formeditassocbuttonsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<FormEditAssocButtonAPI[]>('getFormEditAssocButtons', []))
      );
  }

  /** GET formeditassocbutton by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormEditAssocButtonAPI> {
    return this.getFormEditAssocButton(id, GONG__StackPath, frontRepo)
  }
  getFormEditAssocButton(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormEditAssocButtonAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.formeditassocbuttonsUrl}/${id}`;
    return this.http.get<FormEditAssocButtonAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched formeditassocbutton id=${id}`)),
      catchError(this.handleError<FormEditAssocButtonAPI>(`getFormEditAssocButton id=${id}`))
    );
  }

  // postFront copy formeditassocbutton to a version with encoded pointers and post to the back
  postFront(formeditassocbutton: FormEditAssocButton, GONG__StackPath: string): Observable<FormEditAssocButtonAPI> {
    let formeditassocbuttonAPI = new FormEditAssocButtonAPI
    CopyFormEditAssocButtonToFormEditAssocButtonAPI(formeditassocbutton, formeditassocbuttonAPI)
    const id = typeof formeditassocbuttonAPI === 'number' ? formeditassocbuttonAPI : formeditassocbuttonAPI.ID
    const url = `${this.formeditassocbuttonsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormEditAssocButtonAPI>(url, formeditassocbuttonAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<FormEditAssocButtonAPI>('postFormEditAssocButton'))
    );
  }
  
  /** POST: add a new formeditassocbutton to the server */
  post(formeditassocbuttondb: FormEditAssocButtonAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormEditAssocButtonAPI> {
    return this.postFormEditAssocButton(formeditassocbuttondb, GONG__StackPath, frontRepo)
  }
  postFormEditAssocButton(formeditassocbuttondb: FormEditAssocButtonAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormEditAssocButtonAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormEditAssocButtonAPI>(this.formeditassocbuttonsUrl, formeditassocbuttondb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted formeditassocbuttondb id=${formeditassocbuttondb.ID}`)
      }),
      catchError(this.handleError<FormEditAssocButtonAPI>('postFormEditAssocButton'))
    );
  }

  /** DELETE: delete the formeditassocbuttondb from the server */
  delete(formeditassocbuttondb: FormEditAssocButtonAPI | number, GONG__StackPath: string): Observable<FormEditAssocButtonAPI> {
    return this.deleteFormEditAssocButton(formeditassocbuttondb, GONG__StackPath)
  }
  deleteFormEditAssocButton(formeditassocbuttondb: FormEditAssocButtonAPI | number, GONG__StackPath: string): Observable<FormEditAssocButtonAPI> {
    const id = typeof formeditassocbuttondb === 'number' ? formeditassocbuttondb : formeditassocbuttondb.ID;
    const url = `${this.formeditassocbuttonsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<FormEditAssocButtonAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted formeditassocbuttondb id=${id}`)),
      catchError(this.handleError<FormEditAssocButtonAPI>('deleteFormEditAssocButton'))
    );
  }

  // updateFront copy formeditassocbutton to a version with encoded pointers and update to the back
  updateFront(formeditassocbutton: FormEditAssocButton, GONG__StackPath: string): Observable<FormEditAssocButtonAPI> {
    let formeditassocbuttonAPI = new FormEditAssocButtonAPI
    CopyFormEditAssocButtonToFormEditAssocButtonAPI(formeditassocbutton, formeditassocbuttonAPI)
    const id = typeof formeditassocbuttonAPI === 'number' ? formeditassocbuttonAPI : formeditassocbuttonAPI.ID
    const url = `${this.formeditassocbuttonsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<FormEditAssocButtonAPI>(url, formeditassocbuttonAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<FormEditAssocButtonAPI>('updateFormEditAssocButton'))
    );
  }

  /** PUT: update the formeditassocbuttondb on the server */
  update(formeditassocbuttondb: FormEditAssocButtonAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormEditAssocButtonAPI> {
    return this.updateFormEditAssocButton(formeditassocbuttondb, GONG__StackPath, frontRepo)
  }
  updateFormEditAssocButton(formeditassocbuttondb: FormEditAssocButtonAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormEditAssocButtonAPI> {
    const id = typeof formeditassocbuttondb === 'number' ? formeditassocbuttondb : formeditassocbuttondb.ID;
    const url = `${this.formeditassocbuttonsUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<FormEditAssocButtonAPI>(url, formeditassocbuttondb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated formeditassocbuttondb id=${formeditassocbuttondb.ID}`)
      }),
      catchError(this.handleError<FormEditAssocButtonAPI>('updateFormEditAssocButton'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in FormEditAssocButtonService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("FormEditAssocButtonService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}

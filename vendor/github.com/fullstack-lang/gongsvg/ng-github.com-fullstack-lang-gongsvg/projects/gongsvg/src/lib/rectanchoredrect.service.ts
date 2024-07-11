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

import { RectAnchoredRectAPI } from './rectanchoredrect-api'
import { RectAnchoredRect, CopyRectAnchoredRectToRectAnchoredRectAPI } from './rectanchoredrect'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class RectAnchoredRectService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  RectAnchoredRectServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private rectanchoredrectsUrl: string

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
    this.rectanchoredrectsUrl = origin + '/api/github.com/fullstack-lang/gongsvg/go/v1/rectanchoredrects';
  }

  /** GET rectanchoredrects from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectAnchoredRectAPI[]> {
    return this.getRectAnchoredRects(GONG__StackPath, frontRepo)
  }
  getRectAnchoredRects(GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectAnchoredRectAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<RectAnchoredRectAPI[]>(this.rectanchoredrectsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<RectAnchoredRectAPI[]>('getRectAnchoredRects', []))
      );
  }

  /** GET rectanchoredrect by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectAnchoredRectAPI> {
    return this.getRectAnchoredRect(id, GONG__StackPath, frontRepo)
  }
  getRectAnchoredRect(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectAnchoredRectAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.rectanchoredrectsUrl}/${id}`;
    return this.http.get<RectAnchoredRectAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched rectanchoredrect id=${id}`)),
      catchError(this.handleError<RectAnchoredRectAPI>(`getRectAnchoredRect id=${id}`))
    );
  }

  // postFront copy rectanchoredrect to a version with encoded pointers and post to the back
  postFront(rectanchoredrect: RectAnchoredRect, GONG__StackPath: string): Observable<RectAnchoredRectAPI> {
    let rectanchoredrectAPI = new RectAnchoredRectAPI
    CopyRectAnchoredRectToRectAnchoredRectAPI(rectanchoredrect, rectanchoredrectAPI)
    const id = typeof rectanchoredrectAPI === 'number' ? rectanchoredrectAPI : rectanchoredrectAPI.ID
    const url = `${this.rectanchoredrectsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<RectAnchoredRectAPI>(url, rectanchoredrectAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<RectAnchoredRectAPI>('postRectAnchoredRect'))
    );
  }
  
  /** POST: add a new rectanchoredrect to the server */
  post(rectanchoredrectdb: RectAnchoredRectAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectAnchoredRectAPI> {
    return this.postRectAnchoredRect(rectanchoredrectdb, GONG__StackPath, frontRepo)
  }
  postRectAnchoredRect(rectanchoredrectdb: RectAnchoredRectAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectAnchoredRectAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<RectAnchoredRectAPI>(this.rectanchoredrectsUrl, rectanchoredrectdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted rectanchoredrectdb id=${rectanchoredrectdb.ID}`)
      }),
      catchError(this.handleError<RectAnchoredRectAPI>('postRectAnchoredRect'))
    );
  }

  /** DELETE: delete the rectanchoredrectdb from the server */
  delete(rectanchoredrectdb: RectAnchoredRectAPI | number, GONG__StackPath: string): Observable<RectAnchoredRectAPI> {
    return this.deleteRectAnchoredRect(rectanchoredrectdb, GONG__StackPath)
  }
  deleteRectAnchoredRect(rectanchoredrectdb: RectAnchoredRectAPI | number, GONG__StackPath: string): Observable<RectAnchoredRectAPI> {
    const id = typeof rectanchoredrectdb === 'number' ? rectanchoredrectdb : rectanchoredrectdb.ID;
    const url = `${this.rectanchoredrectsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<RectAnchoredRectAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted rectanchoredrectdb id=${id}`)),
      catchError(this.handleError<RectAnchoredRectAPI>('deleteRectAnchoredRect'))
    );
  }

  // updateFront copy rectanchoredrect to a version with encoded pointers and update to the back
  updateFront(rectanchoredrect: RectAnchoredRect, GONG__StackPath: string): Observable<RectAnchoredRectAPI> {
    let rectanchoredrectAPI = new RectAnchoredRectAPI
    CopyRectAnchoredRectToRectAnchoredRectAPI(rectanchoredrect, rectanchoredrectAPI)
    const id = typeof rectanchoredrectAPI === 'number' ? rectanchoredrectAPI : rectanchoredrectAPI.ID
    const url = `${this.rectanchoredrectsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<RectAnchoredRectAPI>(url, rectanchoredrectAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<RectAnchoredRectAPI>('updateRectAnchoredRect'))
    );
  }

  /** PUT: update the rectanchoredrectdb on the server */
  update(rectanchoredrectdb: RectAnchoredRectAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectAnchoredRectAPI> {
    return this.updateRectAnchoredRect(rectanchoredrectdb, GONG__StackPath, frontRepo)
  }
  updateRectAnchoredRect(rectanchoredrectdb: RectAnchoredRectAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectAnchoredRectAPI> {
    const id = typeof rectanchoredrectdb === 'number' ? rectanchoredrectdb : rectanchoredrectdb.ID;
    const url = `${this.rectanchoredrectsUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<RectAnchoredRectAPI>(url, rectanchoredrectdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated rectanchoredrectdb id=${rectanchoredrectdb.ID}`)
      }),
      catchError(this.handleError<RectAnchoredRectAPI>('updateRectAnchoredRect'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in RectAnchoredRectService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("RectAnchoredRectService" + error); // log to console instead

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

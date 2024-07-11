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

import { GongLinkAPI } from './gonglink-api'
import { GongLink, CopyGongLinkToGongLinkAPI } from './gonglink'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class GongLinkService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  GongLinkServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private gonglinksUrl: string

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
    this.gonglinksUrl = origin + '/api/github.com/fullstack-lang/gong/go/v1/gonglinks';
  }

  /** GET gonglinks from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongLinkAPI[]> {
    return this.getGongLinks(GONG__StackPath, frontRepo)
  }
  getGongLinks(GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongLinkAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<GongLinkAPI[]>(this.gonglinksUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<GongLinkAPI[]>('getGongLinks', []))
      );
  }

  /** GET gonglink by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongLinkAPI> {
    return this.getGongLink(id, GONG__StackPath, frontRepo)
  }
  getGongLink(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongLinkAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.gonglinksUrl}/${id}`;
    return this.http.get<GongLinkAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched gonglink id=${id}`)),
      catchError(this.handleError<GongLinkAPI>(`getGongLink id=${id}`))
    );
  }

  // postFront copy gonglink to a version with encoded pointers and post to the back
  postFront(gonglink: GongLink, GONG__StackPath: string): Observable<GongLinkAPI> {
    let gonglinkAPI = new GongLinkAPI
    CopyGongLinkToGongLinkAPI(gonglink, gonglinkAPI)
    const id = typeof gonglinkAPI === 'number' ? gonglinkAPI : gonglinkAPI.ID
    const url = `${this.gonglinksUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<GongLinkAPI>(url, gonglinkAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<GongLinkAPI>('postGongLink'))
    );
  }
  
  /** POST: add a new gonglink to the server */
  post(gonglinkdb: GongLinkAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongLinkAPI> {
    return this.postGongLink(gonglinkdb, GONG__StackPath, frontRepo)
  }
  postGongLink(gonglinkdb: GongLinkAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongLinkAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<GongLinkAPI>(this.gonglinksUrl, gonglinkdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted gonglinkdb id=${gonglinkdb.ID}`)
      }),
      catchError(this.handleError<GongLinkAPI>('postGongLink'))
    );
  }

  /** DELETE: delete the gonglinkdb from the server */
  delete(gonglinkdb: GongLinkAPI | number, GONG__StackPath: string): Observable<GongLinkAPI> {
    return this.deleteGongLink(gonglinkdb, GONG__StackPath)
  }
  deleteGongLink(gonglinkdb: GongLinkAPI | number, GONG__StackPath: string): Observable<GongLinkAPI> {
    const id = typeof gonglinkdb === 'number' ? gonglinkdb : gonglinkdb.ID;
    const url = `${this.gonglinksUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<GongLinkAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted gonglinkdb id=${id}`)),
      catchError(this.handleError<GongLinkAPI>('deleteGongLink'))
    );
  }

  // updateFront copy gonglink to a version with encoded pointers and update to the back
  updateFront(gonglink: GongLink, GONG__StackPath: string): Observable<GongLinkAPI> {
    let gonglinkAPI = new GongLinkAPI
    CopyGongLinkToGongLinkAPI(gonglink, gonglinkAPI)
    const id = typeof gonglinkAPI === 'number' ? gonglinkAPI : gonglinkAPI.ID
    const url = `${this.gonglinksUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<GongLinkAPI>(url, gonglinkAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<GongLinkAPI>('updateGongLink'))
    );
  }

  /** PUT: update the gonglinkdb on the server */
  update(gonglinkdb: GongLinkAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongLinkAPI> {
    return this.updateGongLink(gonglinkdb, GONG__StackPath, frontRepo)
  }
  updateGongLink(gonglinkdb: GongLinkAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongLinkAPI> {
    const id = typeof gonglinkdb === 'number' ? gonglinkdb : gonglinkdb.ID;
    const url = `${this.gonglinksUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<GongLinkAPI>(url, gonglinkdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated gonglinkdb id=${gonglinkdb.ID}`)
      }),
      catchError(this.handleError<GongLinkAPI>('updateGongLink'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in GongLinkService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("GongLinkService" + error); // log to console instead

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

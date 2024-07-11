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

import { GongStructAPI } from './gongstruct-api'
import { GongStruct, CopyGongStructToGongStructAPI } from './gongstruct'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { GongBasicFieldAPI } from './gongbasicfield-api'
import { GongTimeFieldAPI } from './gongtimefield-api'
import { PointerToGongStructFieldAPI } from './pointertogongstructfield-api'
import { SliceOfPointerToGongStructFieldAPI } from './sliceofpointertogongstructfield-api'

@Injectable({
  providedIn: 'root'
})
export class GongStructService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  GongStructServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private gongstructsUrl: string

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
    this.gongstructsUrl = origin + '/api/github.com/fullstack-lang/gong/go/v1/gongstructs';
  }

  /** GET gongstructs from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongStructAPI[]> {
    return this.getGongStructs(GONG__StackPath, frontRepo)
  }
  getGongStructs(GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongStructAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<GongStructAPI[]>(this.gongstructsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<GongStructAPI[]>('getGongStructs', []))
      );
  }

  /** GET gongstruct by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongStructAPI> {
    return this.getGongStruct(id, GONG__StackPath, frontRepo)
  }
  getGongStruct(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongStructAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.gongstructsUrl}/${id}`;
    return this.http.get<GongStructAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched gongstruct id=${id}`)),
      catchError(this.handleError<GongStructAPI>(`getGongStruct id=${id}`))
    );
  }

  // postFront copy gongstruct to a version with encoded pointers and post to the back
  postFront(gongstruct: GongStruct, GONG__StackPath: string): Observable<GongStructAPI> {
    let gongstructAPI = new GongStructAPI
    CopyGongStructToGongStructAPI(gongstruct, gongstructAPI)
    const id = typeof gongstructAPI === 'number' ? gongstructAPI : gongstructAPI.ID
    const url = `${this.gongstructsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<GongStructAPI>(url, gongstructAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<GongStructAPI>('postGongStruct'))
    );
  }
  
  /** POST: add a new gongstruct to the server */
  post(gongstructdb: GongStructAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongStructAPI> {
    return this.postGongStruct(gongstructdb, GONG__StackPath, frontRepo)
  }
  postGongStruct(gongstructdb: GongStructAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongStructAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<GongStructAPI>(this.gongstructsUrl, gongstructdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted gongstructdb id=${gongstructdb.ID}`)
      }),
      catchError(this.handleError<GongStructAPI>('postGongStruct'))
    );
  }

  /** DELETE: delete the gongstructdb from the server */
  delete(gongstructdb: GongStructAPI | number, GONG__StackPath: string): Observable<GongStructAPI> {
    return this.deleteGongStruct(gongstructdb, GONG__StackPath)
  }
  deleteGongStruct(gongstructdb: GongStructAPI | number, GONG__StackPath: string): Observable<GongStructAPI> {
    const id = typeof gongstructdb === 'number' ? gongstructdb : gongstructdb.ID;
    const url = `${this.gongstructsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<GongStructAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted gongstructdb id=${id}`)),
      catchError(this.handleError<GongStructAPI>('deleteGongStruct'))
    );
  }

  // updateFront copy gongstruct to a version with encoded pointers and update to the back
  updateFront(gongstruct: GongStruct, GONG__StackPath: string): Observable<GongStructAPI> {
    let gongstructAPI = new GongStructAPI
    CopyGongStructToGongStructAPI(gongstruct, gongstructAPI)
    const id = typeof gongstructAPI === 'number' ? gongstructAPI : gongstructAPI.ID
    const url = `${this.gongstructsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<GongStructAPI>(url, gongstructAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<GongStructAPI>('updateGongStruct'))
    );
  }

  /** PUT: update the gongstructdb on the server */
  update(gongstructdb: GongStructAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongStructAPI> {
    return this.updateGongStruct(gongstructdb, GONG__StackPath, frontRepo)
  }
  updateGongStruct(gongstructdb: GongStructAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongStructAPI> {
    const id = typeof gongstructdb === 'number' ? gongstructdb : gongstructdb.ID;
    const url = `${this.gongstructsUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<GongStructAPI>(url, gongstructdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated gongstructdb id=${gongstructdb.ID}`)
      }),
      catchError(this.handleError<GongStructAPI>('updateGongStruct'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in GongStructService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("GongStructService" + error); // log to console instead

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

export namespace service {
	
	export class ConnCollection {
	    ID: number;
	    CollectionName: string;
	
	    static createFrom(source: any = {}) {
	        return new ConnCollection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CollectionName = source["CollectionName"];
	    }
	}
	export class ConnInfo {
	    ID: number;
	    Collection_ID: sql.NullInt16;
	    ConnName: string;
	    IP: string;
	    Port: number;
	    UserName: sql.NullString;
	    Password: sql.NullString;
	    Key: sql.NullString;
	
	    static createFrom(source: any = {}) {
	        return new ConnInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Collection_ID = this.convertValues(source["Collection_ID"], sql.NullInt16);
	        this.ConnName = source["ConnName"];
	        this.IP = source["IP"];
	        this.Port = source["Port"];
	        this.UserName = this.convertValues(source["UserName"], sql.NullString);
	        this.Password = this.convertValues(source["Password"], sql.NullString);
	        this.Key = this.convertValues(source["Key"], sql.NullString);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace sql {
	
	export class NullInt16 {
	    Int16: number;
	    Valid: boolean;
	
	    static createFrom(source: any = {}) {
	        return new NullInt16(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Int16 = source["Int16"];
	        this.Valid = source["Valid"];
	    }
	}
	export class NullString {
	    String: string;
	    Valid: boolean;
	
	    static createFrom(source: any = {}) {
	        return new NullString(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.String = source["String"];
	        this.Valid = source["Valid"];
	    }
	}

}

export namespace types {
	
	export class Response {
	    ResponseCode: number;
	    ResponseMsg: string;
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ResponseCode = source["ResponseCode"];
	        this.ResponseMsg = source["ResponseMsg"];
	    }
	}

}


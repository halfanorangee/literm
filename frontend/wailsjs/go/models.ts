export namespace service {
	
	export class ConnInfo {
	    ID: number;
	    Collection_ID: number;
	    ConnName: string;
	    IP: string;
	    Port: string;
	    UserName: string;
	    Password: string;
	    AuthKey: string;
	    AuthType: string;
	
	    static createFrom(source: any = {}) {
	        return new ConnInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Collection_ID = source["Collection_ID"];
	        this.ConnName = source["ConnName"];
	        this.IP = source["IP"];
	        this.Port = source["Port"];
	        this.UserName = source["UserName"];
	        this.Password = source["Password"];
	        this.AuthKey = source["AuthKey"];
	        this.AuthType = source["AuthType"];
	    }
	}
	export class CollectionConnRel {
	    ID: number;
	    CollectionName: string;
	    ConnInfos: ConnInfo[];
	
	    static createFrom(source: any = {}) {
	        return new CollectionConnRel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CollectionName = source["CollectionName"];
	        this.ConnInfos = this.convertValues(source["ConnInfos"], ConnInfo);
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

export namespace types {
	
	export class ConnInsertInfo {
	    CollectionId: number;
	    Name: string;
	    Comment: string;
	    IP: string;
	    Port: string;
	    UserName: string;
	    Password: string;
	    Authkey: string;
	    AuthType: string;
	
	    static createFrom(source: any = {}) {
	        return new ConnInsertInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.CollectionId = source["CollectionId"];
	        this.Name = source["Name"];
	        this.Comment = source["Comment"];
	        this.IP = source["IP"];
	        this.Port = source["Port"];
	        this.UserName = source["UserName"];
	        this.Password = source["Password"];
	        this.Authkey = source["Authkey"];
	        this.AuthType = source["AuthType"];
	    }
	}
	export class Response {
	    ResponseCode: string;
	    ResponseMsg: string;
	    ResponseInfo: string;
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ResponseCode = source["ResponseCode"];
	        this.ResponseMsg = source["ResponseMsg"];
	        this.ResponseInfo = source["ResponseInfo"];
	    }
	}

}


export namespace internal {
	
	export class ResBody {
	    id: string;
	    content: string;
	    author: string;
	    assign_date: string;
	    ad_url: string;
	    share_url: string;
	    // Go type: struct {}
	    share_urls: any;
	    origin_img_urls: string[];
	    share_img_urls: string[];
	    join_num: number;
	    translation: string;
	    poster_img_urls: string[];
	    // Go type: struct {}
	    track_object: any;
	    daily_audio_urls: string;
	
	    static createFrom(source: any = {}) {
	        return new ResBody(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.content = source["content"];
	        this.author = source["author"];
	        this.assign_date = source["assign_date"];
	        this.ad_url = source["ad_url"];
	        this.share_url = source["share_url"];
	        this.share_urls = this.convertValues(source["share_urls"], Object);
	        this.origin_img_urls = source["origin_img_urls"];
	        this.share_img_urls = source["share_img_urls"];
	        this.join_num = source["join_num"];
	        this.translation = source["translation"];
	        this.poster_img_urls = source["poster_img_urls"];
	        this.track_object = this.convertValues(source["track_object"], Object);
	        this.daily_audio_urls = source["daily_audio_urls"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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


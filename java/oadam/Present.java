package oadam;

import org.codehaus.jackson.annotate.JsonIgnoreProperties;
import org.joda.time.DateTime;

import com.googlecode.objectify.Key;
import com.googlecode.objectify.annotation.Entity;
import com.googlecode.objectify.annotation.Id;
import com.googlecode.objectify.annotation.Parent;

@Entity
@JsonIgnoreProperties({"parent"})
public class Present {
	@Id public Long id;
	@Parent public Key<Party> parent;
	public String title;
	public String description;
	public Long to;
	public Long createdBy;
	public DateTime creationDate;
	public Long offeredBy;
	public DateTime offeredDate;
	/** Presents are sorted by this.
	 *  If it is null creationDate is used instead*/
	public DateTime sortDate;
	public Long deletedBy;
	
	public Present(Long id, Key<Party> parent, String title, String description, Long to,
			Long createdBy, DateTime creationDate, Long offeredBy,
			DateTime offeredDate, DateTime sortDate, Long deletedBy) {
		super();
		this.id = id;
		this.parent = parent;
		this.title = title;
		this.description = description;
		this.to = to;
		this.createdBy = createdBy;
		this.creationDate = creationDate;
		this.offeredBy = offeredBy;
		this.offeredDate = offeredDate;
		this.sortDate = sortDate;
		this.deletedBy = deletedBy;
	}
	
	public Present() {
	}
}

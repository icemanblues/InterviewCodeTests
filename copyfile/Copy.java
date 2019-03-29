/**
 * @author roland
 *
 */
public class Copy {
	private int dataCenterSource;
	private int dataCenterDest;
	private int dataSetId;

	public Copy(int dataCenterSource, int dataCenterDest, int dataSetId) {
		this.dataCenterSource = dataCenterSource;
		this.dataCenterDest = dataCenterDest;
		this.dataSetId = dataSetId;
	}

	public int getDataCenterSource() {
		return dataCenterSource;
	}

	public int getDataCenterDest() {
		return dataCenterDest;
	}

	public int getDataSetId() {
		return dataSetId;
	}

	public void setDataSetId(int dataSetId) {
		this.dataSetId = dataSetId;
	}

	//  <data-set-id> <from> <to>
	@Override
	public String toString() {
		return  dataSetId + " " + dataCenterSource + " " + dataCenterDest;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + dataCenterDest;
		result = prime * result + dataCenterSource;
		result = prime * result + dataSetId;
		return result;
	}

	@Override
	public boolean equals(Object obj) {
		if (this == obj)
			return true;
		if (obj == null)
			return false;
		if (getClass() != obj.getClass())
			return false;
		Copy other = (Copy) obj;
		if (dataCenterDest != other.dataCenterDest)
			return false;
		if (dataCenterSource != other.dataCenterSource)
			return false;
		if (dataSetId != other.dataSetId)
			return false;
		return true;
	}
}

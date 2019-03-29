import java.util.Collections;
import java.util.HashSet;
import java.util.Set;

/**
 * @author <a href="mailto:roland.kluge@gmail.com">Roland Kluge</a>
 *
 */
public class DataCenter {
	private final int id;
	private Set<Integer> dataSets;

	private static Set<Integer> allDataSets = new HashSet<Integer>();


	public DataCenter(int id) {
		this.id = id;
		dataSets = new HashSet<Integer>();
	}

	public void addDataSet(int dataSetId) {
		dataSets.add(dataSetId);
		allDataSets.add(dataSetId);
	}

	private static Set<Integer> getAllDataSets() {
		return Collections.unmodifiableSet(allDataSets);
	}

	public Set<Integer> getMissingDataSets() {
		Set<Integer> missingData = new HashSet<Integer>();
		for(Integer data: allDataSets) {
			if(!dataSets.contains(data)) {
				missingData.add(data);
			}
		}

		return Collections.unmodifiableSet(missingData);
	}

	public int getId() {
		return id;
	}
}

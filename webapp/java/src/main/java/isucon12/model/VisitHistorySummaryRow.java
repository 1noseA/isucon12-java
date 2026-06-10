package isucon12.model;

import java.util.Date;

// 閲覧履歴の集計
public class VisitHistorySummaryRow {
    private String playerId;
    private Date minCreatedAt;

    public String getPlayerId() {
        return playerId;
    }

    public void setPlayerId(String playerId) {
        this.playerId = playerId;
    }

    public Date getMinCreatedAt() {
        return minCreatedAt;
    }

    public void setMinCreatedAt(Date minCreatedAt) {
        this.minCreatedAt = minCreatedAt;
    }
}
